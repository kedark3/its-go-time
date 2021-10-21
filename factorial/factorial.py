def factorial(number):
    answer = 1
    for i in range(1,number+2):
        answer *= i
    return answer

for i in range(5):
    print(factorial(i))