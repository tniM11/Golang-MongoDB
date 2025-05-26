db = db.getSiblingDB('todos');
db.createCollection('todos');
db.todos.insertOne({ title: 'Welcome', description: 'Sample todo item' });

