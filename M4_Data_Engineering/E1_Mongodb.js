// Part 1: Basic MongoDB Commands and Queries

// 1. Create Collections and Insert Data
db.customers.insertMany([
  {
    _id: ObjectId(),
    name: "John Doe",
    email: "johndoe@example.com",
    address: { street: "123 Main St", city: "Springfield", zipcode: "12345" },
    phone: "555-1234",
    registration_date: ISODate("2023-01-01T12:00:00Z"),
  },
  {
    _id: ObjectId(),
    name: "Jane Smith",
    email: "janesmith@example.com",
    address: { street: "456 Elm St", city: "Greendale", zipcode: "67890" },
    phone: "555-5678",
    registration_date: ISODate("2023-02-01T12:00:00Z"),
  },
  {
    _id: ObjectId(),
    name: "Bob Brown",
    email: "bobbrown@example.com",
    address: { street: "789 Oak St", city: "Shelbyville", zipcode: "13579" },
    phone: "555-8765",
    registration_date: ISODate("2023-03-01T12:00:00Z"),
  },
  {
    _id: ObjectId(),
    name: "Alice White",
    email: "alicewhite@example.com",
    address: { street: "101 Pine St", city: "Capital City", zipcode: "24680" },
    phone: "555-4321",
    registration_date: ISODate("2023-04-01T12:00:00Z"),
  },
  {
    _id: ObjectId(),
    name: "Charlie Black",
    email: "charlieblack@example.com",
    address: { street: "202 Maple St", city: "Ogdenville", zipcode: "11223" },
    phone: "555-8765",
    registration_date: ISODate("2023-05-01T12:00:00Z"),
  },
]);

// Insert orders with references to customer _id
db.orders.insertMany([
  {
    _id: ObjectId(),
    order_id: "ORD123456",
    customer_id: ObjectId("customerId1"),
    order_date: ISODate("2023-05-15T14:00:00Z"),
    status: "shipped",
    items: [
      { product_name: "Laptop", quantity: 1, price: 1500 },
      { product_name: "Mouse", quantity: 2, price: 25 },
    ],
    total_value: 1550,
  },
  {
    _id: ObjectId(),
    order_id: "ORD123457",
    customer_id: ObjectId("customerId2"),
    order_date: ISODate("2023-06-10T10:30:00Z"),
    status: "delivered",
    items: [{ product_name: "Tablet", quantity: 1, price: 800 }],
    total_value: 800,
  },
  {
    _id: ObjectId(),
    order_id: "ORD123458",
    customer_id: ObjectId("customerId3"),
    order_date: ISODate("2023-06-12T12:45:00Z"),
    status: "shipped",
    items: [{ product_name: "Phone", quantity: 1, price: 700 }],
    total_value: 700,
  },
  {
    _id: ObjectId(),
    order_id: "ORD123459",
    customer_id: ObjectId("customerId4"),
    order_date: ISODate("2023-06-13T15:20:00Z"),
    status: "processing",
    items: [{ product_name: "Monitor", quantity: 2, price: 300 }],
    total_value: 600,
  },
  {
    _id: ObjectId(),
    order_id: "ORD123460",
    customer_id: ObjectId("customerId5"),
    order_date: ISODate("2023-06-15T16:05:00Z"),
    status: "shipped",
    items: [{ product_name: "Keyboard", quantity: 1, price: 100 }],
    total_value: 100,
  },
]);

// 2. Find Orders for a Specific Customer
const customer = db.customers.findOne({ name: "John Doe" });
db.orders.find({ customer_id: customer._id });

// 3. Find Customer for a Specific Order
const order = db.orders.findOne({ order_id: "ORD123456" });
db.customers.findOne({ _id: order.customer_id });

// 4. Update Order Status
db.orders.updateOne(
  { order_id: "ORD123456" },
  { $set: { status: "delivered" } }
);

// 5. Delete an Order
db.orders.deleteOne({ order_id: "ORD123456" });

// Part 2: Aggregation Pipeline

// 1. Calculate Total Value of All Orders by Customer
db.orders.aggregate([
  { $group: { _id: "$customer_id", totalValue: { $sum: "$total_value" } } },
  {
    $lookup: {
      from: "customers",
      localField: "_id",
      foreignField: "_id",
      as: "customerInfo",
    },
  },
  { $unwind: "$customerInfo" },
  { $project: { "customerInfo.name": 1, totalValue: 1 } },
]);

// 2. Group Orders by Status
db.orders.aggregate([{ $group: { _id: "$status", count: { $sum: 1 } } }]);

// 3. List Customers with Their Recent Orders
db.orders.aggregate([
  { $sort: { order_date: -1 } },
  { $group: { _id: "$customer_id", mostRecentOrder: { $first: "$$ROOT" } } },
  {
    $lookup: {
      from: "customers",
      localField: "_id",
      foreignField: "_id",
      as: "customerInfo",
    },
  },
  { $unwind: "$customerInfo" },
  {
    $project: {
      "customerInfo.name": 1,
      "customerInfo.email": 1,
      mostRecentOrder: 1,
    },
  },
]);

// 4. Find the Most Expensive Order by Customer
db.orders.aggregate([
  { $sort: { total_value: -1 } },
  { $group: { _id: "$customer_id", mostExpensiveOrder: { $first: "$$ROOT" } } },
  {
    $lookup: {
      from: "customers",
      localField: "_id",
      foreignField: "_id",
      as: "customerInfo",
    },
  },
  { $unwind: "$customerInfo" },
  { $project: { "customerInfo.name": 1, mostExpensiveOrder: 1 } },
]);

// Part 3: Real-World Scenario with Relationships

// 1. Find All Customers Who Placed Orders in the Last Month
db.orders.aggregate([
  {
    $match: {
      order_date: { $gte: new Date(new Date() - 30 * 24 * 60 * 60 * 1000) },
    },
  },
  { $group: { _id: "$customer_id", recentOrder: { $max: "$order_date" } } },
  {
    $lookup: {
      from: "customers",
      localField: "_id",
      foreignField: "_id",
      as: "customerInfo",
    },
  },
  { $unwind: "$customerInfo" },
  {
    $project: {
      "customerInfo.name": 1,
      "customerInfo.email": 1,
      recentOrder: 1,
    },
  },
]);

// 2. Find All Products Ordered by a Specific Customer
db.orders.aggregate([
  { $match: { customer_id: customer._id } },
  { $unwind: "$items" },
  {
    $group: {
      _id: "$items.product_name",
      totalQuantity: { $sum: "$items.quantity" },
    },
  },
]);

// 3. Find the Top 3 Customers with the Most Expensive Total Orders
db.orders.aggregate([
  { $group: { _id: "$customer_id", totalSpent: { $sum: "$total_value" } } },
  { $sort: { totalSpent: -1 } },
  { $limit: 3 },
  {
    $lookup: {
      from: "customers",
      localField: "_id",
      foreignField: "_id",
      as: "customerInfo",
    },
  },
  { $unwind: "$customerInfo" },
  { $project: { "customerInfo.name": 1, totalSpent: 1 } },
]);

// 4. Add a New Order for an Existing Customer
const jane = db.customers.findOne({ name: "Jane Smith" });
db.orders.insertOne({
  order_id: "ORD123461",
  customer_id: jane._id,
  order_date: new Date(),
  status: "processing",
  items: [
    { product_name: "Smartphone", quantity: 1, price: 600 },
    { product_name: "Headphones", quantity: 2, price: 50 },
  ],
  total_value: 700,
});

// Part 4: Bonus Challenge

// 1. Find Customers Who Have Not Placed Orders
db.customers.aggregate([
  {
    $lookup: {
      from: "orders",
      localField: "_id",
      foreignField: "customer_id",
      as: "orders",
    },
  },
  { $match: { "orders.0": { $exists: false } } },
  { $project: { name: 1, email: 1 } },
]);

// 2. Calculate the Average Number of Items Ordered per Order
db.orders.aggregate([
  { $unwind: "$items" },
  { $group: { _id: "$_id", itemCount: { $sum: 1 } } },
  { $group: { _id: null, avgItems: { $avg: "$itemCount" } } },
]);

// 3. Join Customer and Order Data Using $lookup
db.orders.aggregate([
  {
    $lookup: {
      from: "customers",
      localField: "customer_id",
      foreignField: "_id",
      as: "customerInfo",
    },
  },
  { $unwind: "$customerInfo" },
  {
    $project: {
      "customerInfo.name": 1,
      "customerInfo.email": 1,
      order_id: 1,
      total_value: 1,
      order_date: 1,
    },
  },
]);
