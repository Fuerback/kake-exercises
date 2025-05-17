# Exercise 3

The Kake bakery is starting to sell cakes with different flavors, sizes, frostings. In order to create a menu, the bakery wants to create a list of all possible combinations of cakes given the available characteristics.

The Kake bakery wants to create a program that will take a JSON input and output a list of all possible combinations of cakes.

Input example:
```json
{
  "bakeryName": "Kake",
  "characteristics": [
    {
      "name": "Flavor",
      "values": ["Chocolate", "Vanilla", "Strawberry", "Lemon"]
    },
    {
      "name": "Size",
      "values": ["Small", "Medium", "Large", "Extra Large"]
    },
    {
      "name": "Frosting",
      "values": ["Buttercream", "Fondant", "Whipped Cream", "No Frosting"]
    }
  ]
}

```

Output format example:
```json
[
    {
      "name": "Kake #0",
      "characteristics": [
        {
          "name": "Flavor",
          "value": "Chocolate"
        },
        {
          "name": "Size",
          "value": "Small"
        },
        {
          "name": "Frosting",
          "value": "Buttercream"
        }
      ]
    },
    {
      "name": "Kake #1",
      "characteristics": [
        {
          "name": "Flavor",
          "value": "Vanilla"
        },
        {
          "name": "Size",
          "value": "Medium"
        },
        {
          "name": "Frosting",
          "value": "Fondant"
        }
      ]
    },
    {
      "name": "Kake #2",
      "characteristics": [
        {
          "name": "Flavor",
          "value": "Strawberry"
        },
        {
          "name": "Size",
          "value": "Large"
        },
        {
          "name": "Frosting",
          "value": "Whipped Cream"
        }
      ]
    }
    ...
]
```

## How to run the program:

```bash
go run main.go input.json
```
        