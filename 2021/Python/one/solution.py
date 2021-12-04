def get_input(file_name):
	with open(file_name) as f:
		return list(map(int, f.read().splitlines()))

inputs = get_input('input.txt')

def part_one():
	counter = 0

	for idx in range(1, len(inputs)):
		previous = inputs[idx-1]
		current = inputs[idx]
		if current > previous:
			counter += 1

	return counter

def part_two():
	counter = 0

	for idx in range(3, len(inputs)):
		previous = inputs[idx-3] + inputs[idx-2] + inputs[idx-1]
		current = inputs[idx-2] + inputs[idx-1] + inputs[idx]
		if current > previous:
			counter += 1

	return counter

print('part 1:', part_one())
print('part 2:', part_two())
