# for each number in test.txt, print the amount of occurrences of that number in test.txt and sort the numbers in ascending order

# open file
file = open("test.txt", "r")

# create empty list
list = []

# for each line in file, add the line to the list
for line in file:
    list.append(line)

# sort the list
list.sort()

# create empty dictionary
dict = {}

# for each number in the list, add the number to the dictionary and count the amount of occurrences
for number in list:
    if number in dict:
        dict[number] += 1
    else:
        dict[number] = 1
        
# print the number and the amount of occurrences
for number in dict:
    print(number, dict[number])

# close file
file.close()

# Output:
