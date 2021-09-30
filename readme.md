# coffeemachine

## Entities
```
* Machine: Main intractions component

* Refiller: The only component which is responsible for refilling the store

* Outlet: Outlet which will be commanded to pour a drink. Outlet commands blender to prepare drink

* Blender: Blender is the component which handles a list of drinks maker i.e. greetea maker, blackcoffee maker etc.

* {Drink(Tea)(HotMilk) etc.}-Makers : does the work of making the drink.

* Recipe: a recipe object for a Drink

* IngredientsStore: The main backbone where total capacity is defined and managed

* JsonParser package: Reading the file provided in the assignment and prepare objects out of which differeent machine components are initilised.
```

##Algorithm

```
1) In main function I call BuildDependencies function to read json file.

2) The initialise the MachineObject

3) Machine accesses refiller to refill the store and has a list of outlets.

4) When prepare method of Machine is invoked it requires which outlet of machine User is requesting to get the drink out of.

5) Machine asks Outlet if it could pour the drink requested by user

6) Outlet then calls blender, blender talks to the corresponding DrinkMaker if it supports else it error out.

7) DrinkMaker talks to Ingredients store to get access to ingredients 
    * if not available store returns error.
    * if available store decreases the volume of ingredients it has and return true indicating it has fulfiled the request.

8) Store uses in-memory map to handle Ingredients; it uses sync.Mutex to maintain atomicity
```
