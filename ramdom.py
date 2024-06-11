import random

# Open "data.txt" file in write mode
with open("data.txt", "w") as f:
    # Generate 1000 random numbers between 1 and 999
    for _ in range(1000):
        random_number = random.randint(1, 999)
        # Write each random number on a separate line
        f.write(str(random_number) + "\n")
# Print success message
print("Random numbers generated and saved to data.txt")
