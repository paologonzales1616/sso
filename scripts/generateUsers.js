const { faker } = require("@faker-js/faker");
const fs = require("fs");
const path = require("path");

const user_count = 5;

const users = [...Array(user_count)].map(() => ({
  userId: faker.string.uuid(),
  email: faker.internet.email().toLowerCase(),
  password: faker.internet.password(),
}));

fs.writeFileSync(
  path.resolve(__dirname, "../src/fixtures/users.json"),
  JSON.stringify(users, null, 2)
);
