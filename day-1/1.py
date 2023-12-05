import re

r = 0

with open("calibration.txt") as file:
    lines = file.readlines()

    for line in lines:
        digits = re.findall(r"\d", line)
        r += int(digits[0] + digits[-1])

print(r)

