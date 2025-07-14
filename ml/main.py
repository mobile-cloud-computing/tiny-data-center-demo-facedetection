from typing import Any, Dict, List
from PIL import Image, ImageDraw
from deepface import DeepFace
from fastapi import FastAPI
import shutil
from minio import Minio
import os
import time
import uuid

STORAGE_ACESS_KEY=os.environ.get("STORAGE_ACESS_KEY", "minioadmin")
STORAGE_SECRET_KEY=os.environ.get("STORAGE_SECRET_KEY", "minioadmin")
STORAGE_DNS=os.environ.get("STORAGE_DNS", "localhost:9000")

# Wait for MinIO to start up (not ideal for production!)
time.sleep(5)

client = Minio(STORAGE_DNS,
    secure=False,
    access_key=STORAGE_ACESS_KEY,
    secret_key=STORAGE_SECRET_KEY,
)

app = FastAPI()

@app.get("/test-minio")
def test_minio():
    try:
        buckets = client.list_buckets()
        return {"buckets": [b.name for b in buckets]}
    except Exception as e:
        return {"error": str(e)}

# 1. Expose a post method for getting the faces
#     Input:  A path on minio with the original picture
#             A output minio dir
    
#     Output:
#             200 Okey done 
#             5xx we fail miserable
@app.post("/ml/faces")
def extract_faces(base_image_path: str, output_path: str):
#   1. Create temp folder for this batch
#     /temp/uuid/
#         source.jpg
#         output/
#   2. download the base image and name it source
#   3. found faces and save them in the output folder
#   4. upload all from the output folder into the output minio path

    workDir = getWorkingDir()
    sourcePath = downloadBaseImage(base_image_path, workDir)
    faces = detectFaces(sourcePath)
    metadata = saveFaces(faces, sourcePath,  f"{workDir}/output")
    metadata = uploadOutput(metadata, output_path)
    clean(workDir)
    return castMetadata(metadata)
    

def castMetadata(metadata: Dict) -> Dict:
    result = []

    for file in metadata:
        theResult = {
            "imagePath": file,
            "generalImage": False,
            "confidence": 0
        }

        if "generalImae" in metadata[file]:
            theResult["generalImage"] = metadata[file]["generalImae"]
        
        if "confidence" in metadata[file]:
            theResult["confidence"] = metadata[file]["confidence"]
        
        result.append(theResult)
    
    return result

def clean(workDir:str):
    shutil.rmtree(workDir)

def getWorkingDir()->str:
    batch_id = uuid.uuid4()
    workDir = "./tmp/{batch_id}".format(batch_id=batch_id)
    os.makedirs("{workDir}/output".format(workDir=workDir))
    return workDir

def downloadBaseImage(base_image_path: str,workDir: str) -> str:
    dir = base_image_path.split("/", 1)
    imageType = dir[1].split(".")[1]
    r = client.get_object(dir[0], dir[1])

    print(r.status)
    print(r.headers)

    # TODO - How can I handle multiple tiples of images ?
    sourceImagePath = '{workDir}/source.{imageType}'.format(workDir=workDir,imageType=imageType)
    with open(sourceImagePath, 'wb') as f:
        f.write(r.data)
    print("Image saved successfully!")

    return sourceImagePath

def detectFaces(source:str) -> List[Dict[str, Any]]:
    start_time = time.time()
    face_objs = DeepFace.extract_faces(
        img_path = source
    )
    print("--- Extracting face: %s seconds ---" % (time.time() - start_time))

    return face_objs

def saveFaces(faces: List[Dict[str, Any]], sourcePath:str,  workDir:str):
    metadata = {}

    image = Image.open(sourcePath)
    draw = ImageDraw.Draw(image)

    ## Creating a rectangle on every face
    general_image_path = "{workDir}/output_image.jpg".format(workDir=workDir)
    for i in faces:
        rect_coords = (i["facial_area"]["x"], i["facial_area"]["y"], i["facial_area"]["x"] + i["facial_area"]["w"], i["facial_area"]["y"] + i["facial_area"]["h"])
        draw.rectangle(rect_coords, outline="red", width=3)
    image.save(general_image_path)

    metadata[general_image_path] = {
            "generalImae": True
    }

    ## Saving each single face
    for index, face_obj in enumerate(faces):
        face_image = Image.fromarray((face_obj["face"] * 255).astype('uint8'))
        iamgePath = f"{workDir}/face_{index}.jpg"
        face_image.save(iamgePath)
        metadata[iamgePath] = {
            "confidence": face_obj["confidence"]
        }
    
    return metadata

def uploadOutput(metadata:dict, output_path:str):
    dir = output_path.split("/", 1)
    bucket = dir[0]
    path = dir[1]
    meta = {}
    for file in metadata:
        rawFile = file.split("/")
        rawFile = rawFile[len(rawFile) -1]
        key = f"{path}/output/{rawFile}"
        
        client.fput_object(
            bucket, key, file,
        )
        meta[f"{bucket}/{key}"] = metadata[file]

    return meta

