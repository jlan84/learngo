x = ["python", "go", "java"]
print(id(x))
x[0] = "c++"
print(id(x))
x[0] = 5
print(id(x))
y = "hello"
print(id(y))
y = 5
print(id(y))
print(x)
print(y)