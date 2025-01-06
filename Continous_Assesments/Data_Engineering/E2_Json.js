// Graded Assessment: Working with JSON Data

// Problem:

// You are tasked with implementing a product management system. The system will use JSON data for storing information about products. Each product has the following properties:

// · id: Unique identifier for the product.

// · name: Name of the product.

// · category: Category of the product.

// · price: Price of the product.

// · available: Boolean indicating if the product is in stock.

// The tasks below involve reading JSON data, adding new products, updating product information, and filtering products based on certain conditions.

// Tasks:

// 1. Parse the JSON data:

// Write a function that reads the JSON data (in the format above) and converts it into a usable data structure. You will need to parse the JSON into a JavaScript object.

// 2. Add a new product:

// Write a function to add a new product to the catalog. This product will have the same structure as the others and should be appended to the products array.

// 3. Update the price of a product:

// Write a function that takes a product ID and a new price and updates the price of the product with the given ID. If the product doesn’t exist, the function should return an error message.

// 4. Filter products based on availability:

// Write a function that returns only the products that are available (i.e., available: true).

// 5. Filter products by category:

// Write a function that takes a category name (e.g., "Electronics") and returns all products in that category.

// Initial JSON data (for demonstration purposes)
const jsonData = `[
    { "id": 1, "name": "Laptop", "category": "Electronics", "price": 1500, "available": true },
    { "id": 2, "name": "Chair", "category": "Furniture", "price": 85, "available": false },
    { "id": 3, "name": "Smartphone", "category": "Electronics", "price": 800, "available": true },
    { "id": 4, "name": "Desk", "category": "Furniture", "price": 200, "available": true }
]`;

// 1. Parse the JSON data
function parseJSONData(data) {
  try {
    return JSON.parse(data);
  } catch (error) {
    console.error("Error parsing JSON:", error);
    return [];
  }
}

let products = parseJSONData(jsonData);

// 2. Add a new product
function addProduct(newProduct) {
  // Ensure the new product has all required properties
  const requiredFields = ["id", "name", "category", "price", "available"];
  const hasAllFields = requiredFields.every((field) =>
    newProduct.hasOwnProperty(field)
  );

  if (!hasAllFields) {
    console.error(
      "New product must contain all required fields:",
      requiredFields
    );
    return;
  }

  products.push(newProduct);
}

// 3. Update the price of a product
function updateProductPrice(productId, newPrice) {
  const product = products.find((p) => p.id === productId);
  if (product) {
    product.price = newPrice;
  } else {
    console.error(`Product with ID ${productId} not found.`);
  }
}

// 4. Filter products based on availability
function getAvailableProducts() {
  return products.filter((product) => product.available);
}

// 5. Filter products by category
function getProductsByCategory(category) {
  return products.filter((product) => product.category === category);
}

// Example usage:

// Adding a new product
addProduct({
  id: 5,
  name: "Table",
  category: "Furniture",
  price: 150,
  available: true,
});

// Updating the price of a product
updateProductPrice(3, 750);

// Filtering available products
console.log("Available Products:", getAvailableProducts());

// Filtering products by category "Electronics"
console.log(
  "Electronics Category Products:",
  getProductsByCategory("Electronics")
);
// Initial JSON data (for demonstration purposes)
const jsonData = `[
    { "id": 1, "name": "Laptop", "category": "Electronics", "price": 1500, "available": true },
    { "id": 2, "name": "Chair", "category": "Furniture", "price": 85, "available": false },
    { "id": 3, "name": "Smartphone", "category": "Electronics", "price": 800, "available": true },
    { "id": 4, "name": "Desk", "category": "Furniture", "price": 200, "available": true }
]`;

// 1. Parse the JSON data
function parseJSONData(data) {
  try {
    return JSON.parse(data);
  } catch (error) {
    console.error("Error parsing JSON:", error);
    return [];
  }
}

let products = parseJSONData(jsonData);

// 2. Add a new product
function addProduct(newProduct) {
  // Ensure the new product has all required properties
  const requiredFields = ["id", "name", "category", "price", "available"];
  const hasAllFields = requiredFields.every((field) =>
    newProduct.hasOwnProperty(field)
  );

  if (!hasAllFields) {
    console.error(
      "New product must contain all required fields:",
      requiredFields
    );
    return;
  }

  products.push(newProduct);
}

// 3. Update the price of a product
function updateProductPrice(productId, newPrice) {
  const product = products.find((p) => p.id === productId);
  if (product) {
    product.price = newPrice;
  } else {
    console.error(`Product with ID ${productId} not found.`);
  }
}

// 4. Filter products based on availability
function getAvailableProducts() {
  return products.filter((product) => product.available);
}

// 5. Filter products by category
function getProductsByCategory(category) {
  return products.filter((product) => product.category === category);
}

// Example usage:

// Adding a new product
addProduct({
  id: 5,
  name: "Table",
  category: "Furniture",
  price: 150,
  available: true,
});

// Updating the price of a product
updateProductPrice(3, 750);

// Filtering available products
console.log("Available Products:", getAvailableProducts());

// Filtering products by category "Electronics"
console.log(
  "Electronics Category Products:",
  getProductsByCategory("Electronics")
);
