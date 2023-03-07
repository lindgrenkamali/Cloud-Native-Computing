import sys
from termcolor import colored, cprint

def boards():
    board = []

    for i in range(8):
        
        if(i == 1 or i == 6):
            row = ["b" for i in range(8)]
            board.append(row)

        else:
            row = [" " for i in range(8)]
            board.append(row)
    
    
    for indexY, row in enumerate(board):
        for indexX in row:
            if(indexY == 1):
                cprint("b", "white", end="")

            elif(indexY == 6):
                cprint("b", color="white", on_color="on_white")


def subMarine1():
    with open("files/day1input.txt") as f:
        lines = f.readlines()

        count = 0
        lastValue = int(lines[0])

        for l in lines[1:]:
            currentValue = int(l)
            if(currentValue > lastValue):
                count += 1
                print(currentValue," - större än ", lastValue, " - ANTAL=", count)

            elif(lastValue > currentValue):
                print(currentValue, " - mindre än ", lastValue)

            lastValue = currentValue



def subMarine2():
    with open("files/day2input.txt") as f:
        lines = f.readlines()
       


class Submarine:
    def __init__(self):
        self.x = 0
        self.y = 0
    
    def PrintPosition(self):
        print("Depth:", self.y, "Location:", self.x)
    
    def Run(self):
        with open("files/day2input.txt") as f:
            lines = f.readlines()
            
            for l in lines:
                text = l.split(" ")
                word = text[0]
                number = int(text[1])

                if(word == "forward"):
                    self.x += number

                elif(word == "up"):
                    self.y -= number 

                else:
                    self.y += number
            
                self.PrintPosition()

s = Submarine()
s.Run()