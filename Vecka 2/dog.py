import json
import os


class Dog:
    def __init__(self, name, race, age, weight):
        self.name = name
        self.race = race
        self.age = age
        self.weight = weight
        self.tailLength = self.getTailLength()
            
    def to_dict(self):
        return {"name": self.name, "race": self.race, "age": self.age, "weight":self.weight,
                "tailLength": self.tailLength}

    def getTailLength(dog):

        if(dog.race == "Tax"):
           return 3.7

        else:
            return float(int(dog.age)*int(dog.weight)/10)
    
    def __str__(self):
        
        return("Namn: {}, Ras: {}, Ålder(år): {}, Vikt(kg): {}, Svanslängd(dm): {}".format(self.name, self.race, self.age, self.weight, self.tailLength))

   


class Kennel:
    def __init__(self):
        self.dogs = []

    def addDog(self):
        name = input("What's your dog's name:")
        race = input("What's your dog's race:")
        age = int(input("What's your dog's age:"))
        weight = float(input("What's your dog's weight:"))

        newDog = Dog(name, race, age, weight)
        self.dogs.append(newDog)
        print("{} was added".format(newDog))
    


    def printDogs(self):
        
        number = int(input("Min tail length: "))
        
        for d in self.dogs:
            if(d.tailLength >= number):
                print(d)

    def removeDog(self):
      name = input("Name of the dog to remove:")
      found = False
      for d in self.dogs:
          if(d.name == name):
              self.dogs.remove(d)
              
              found = True
              break
          

      if(found):
            print("{} was removed".format(name))

      else:          
            print("No dog with name {} exists".format(name))       



class System:

    def __init__(self):
        self.kennel = Kennel()
        self.running = True

    def run(self):

        while(self.running):
            print("1: Add dog")
            print("2: List dogs")
            print("3: Remove dog")
            print("4: Exit system")

            number = int(input())

            if(number == 1):
                self.kennel.addDog()
                self.save()

            elif(number == 2):
                self.kennel.printDogs()

            elif(number == 3):
                self.kennel.removeDog()
                self.save()

            elif(number == 4):
                self.running = False

    def save(self):
        with open("kennel.json", "w") as file:
            dogDict = [dog.to_dict() for dog in self.kennel.dogs]
            file.write(json.dumps({"dogs": dogDict}))

    def read(self):
        
            with open("kennel.json", "r") as file:
                kennel = json.load(file)
                
                for d in kennel["dogs"]:
                    self.kennel.dogs.append(Dog(d["name"], d["race"], int(d["age"]), int(d["weight"])))
        
        

s = System()
s.read()

s.run()