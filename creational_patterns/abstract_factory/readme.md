## Abstract Factory Design Pattern

**Abstract Factory** is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

### Problem
You are creating a **furniture shop simulator**. Your code consists of classes that represent:
- A family of related products, say: `Chair` + `Sofa` + `CoffeeTable`.
- Several variants of this family. For example, products `Chair` + `Sofa` + `CoffeeTable` are available in these variants: `Modern`, `Victorian`, `ArtDeco`.

So, you need a way to create individual furniture objects so that they match other objects of the same family. Also, you don’t want to change existing code when adding new products or families of products to the program.


### Solution
1. Declare **abstract product interfaces** for all product types. Then make all **concrete product classes** implement these interfaces.
2. Declare the **abstract factory interface** with a set of creation methods for all abstract products.
3. Implement a set of **concrete factory classes**, one for each product variant.
4. Create **factory initialization** code. It should instantiate one of the concrete factory classes, depending on the application configuration or the current environment.
5. Use this **factory object** for the required operations.


---

#### Reference:
https://refactoring.guru/design-patterns/abstract-factory