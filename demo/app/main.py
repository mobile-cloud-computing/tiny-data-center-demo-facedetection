from deepface import DeepFace
from PIL import Image, ImageDraw

IMAGE = "test2.jpg"
face_objs = DeepFace.extract_faces(
  img_path = IMAGE
)

# Load your image
image = Image.open(IMAGE)

# Create a drawing object
draw = ImageDraw.Draw(image)

# Rectangle coordinates
for i in face_objs:
  rect_coords = (i["facial_area"]["x"], i["facial_area"]["y"], i["facial_area"]["x"] + i["facial_area"]["w"], i["facial_area"]["y"] + i["facial_area"]["h"])
  draw.rectangle(rect_coords, outline="red", width=3)

# Save or display the image
image.save("output_image.jpg")
# image.show()

for index, face_obj in enumerate(face_objs):
  face_image = Image.fromarray((face_obj["face"] * 255).astype('uint8'))
  face_image.save(f"face_{index}.jpg")