def get_input(file_name):
	with open(file_name) as f:
		return f.read().splitlines()

inputs = get_input('input.txt')

# example
# inputs = '''forward 5
# down 5
# forward 8
# up 3
# down 8
# forward 2
# '''.splitlines()

FORWARD = 'forward'
DOWN = 'down'
UP = 'up'

def part_one():
	horizontal_position = 0
	vertical_position = 0

	for value in inputs:
		moviment, qty = value.split(' ')
		qty = int(qty)
		if moviment == FORWARD:
			horizontal_position += qty
		elif moviment == DOWN:
			vertical_position += qty
		elif moviment == UP:
			vertical_position -= qty

	return vertical_position * horizontal_position

def part_two():
	horizontal_position = 0
	vertical_position = 0
	aim = 0

	for value in inputs:
		moviment, qty = value.split(' ')
		qty = int(qty)
		if moviment == FORWARD:
			horizontal_position += qty
			vertical_position += aim * qty
		elif moviment == DOWN:
			aim += qty
		elif moviment == UP:
			aim -= qty

	return vertical_position * horizontal_position

print('part one:', part_one())
print('part two:', part_two())