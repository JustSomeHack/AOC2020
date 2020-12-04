
def getArea():
  file = open('./input.txt', 'r')
  lines = file.readlines()
  file.close()

  area = []

  for line in lines:
    row = []
    for c in line:
      if c == " " or c == "\n":
        break
      row.append(c)
    area.append(row)

  return area

def traverseArea(area, x, y):
  currentX = x
  currentY = y

  trees = 0

  while currentY < len(area):
    if area[currentY][currentX % len(area[currentY])] == '#':
      trees += 1
    currentX += x
    currentY += y
  
  return trees

area = getArea()

a = traverseArea(area, 1 ,1)
b = traverseArea(area, 3 ,1)
c = traverseArea(area, 5 ,1)
d = traverseArea(area, 7 ,1)
e = traverseArea(area, 1 ,2)

print(a*b*c*d*e)
