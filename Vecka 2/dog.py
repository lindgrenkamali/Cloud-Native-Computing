class Dog:
    def __init__(self, name, race, age, weight):
        self.name = name
        self.race = race
        self.age = age
        self.weight = weight
        self.tailLength = self.getTailLength(self)
            

    def getTailLength(dog):

        if(dog.race == "Tax"):
           return 3.7

        else:
            return dog.age*dog.weight/10
    
    def __repr__(self):
        
        return("Namn: %s, Ras: %s, Ålder(år): %d, Vikt(kg): %d, Svanslängd(dm): {0:f}", self.name, self.race, self.age, self.weight,  )

   


class Kennel:
    def __init__(self):
        self.dogs = []

    def addDog(self):
        name = input("What's your dog's name:")
        race = input("What's your dog's race:")
        age = input("What's your dog's age:")
        weight = input("What's your dog's weight:")

        newDog = Dog(name, race, age, weight)
        self.dogs.append(newDog)
        print(Dog, "was added")
    

    def printDogs(self):
        number = input("Min tail length: ")
        
        for d in self.dogs:
            if(d.tailLength >= number):
                print(d)

    def removeDog(self):
      name = input("Name of the dog to remove:")

      for d in self.dogs:
          if(d.name == name):
              self.dogs.remove(d)
              print("%s was removed", name)
              break

      print("No dog with name %s exists", name)        


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

            number = input()

            if(number == 1):
                self.kennel.addDog()

            elif(number == 2):
                self.kennel.printDogs()

            elif(number == 3):
                self.kennel.removeDog()

            elif(number == 4):
                self.running = False