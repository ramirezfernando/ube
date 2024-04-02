class Person {
    constructor(firstName, lastName, age) {
      this.firstName = firstName;
      this.lastName = lastName;
      this.age = age;
    }
  
    getFullName() {
      return `${this.firstName} ${this.lastName}`;
    }
  
    greet() {
      return `Hello, my name is ${this.getFullName()} and I'm ${this.age} years old.`;
    }
  
    celebrateBirthday() {
      this.age++;
      return `Happy birthday ${this.getFullName()}! You are now ${this.age} years old.`;
    }
}
  
const person1 = new Person("John", "Doe", 30);
console.log(person1.greet());
console.log(person1.celebrateBirthday());
