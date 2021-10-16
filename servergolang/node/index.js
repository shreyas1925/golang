const express = require("express");
const app = express();
const PORT = process.env.PORT || 5000;

app.get("/", (req, res) => {
  res.send("Hello from the home side");
});

app.get("/about", (req, res) => {
  res.send("Hello from the about side");
});

app.get("/contact", (req, res) => {
  res.send("Hello from the contact side");
});

app.get("/", (req, res) => {
  res.send("Hello from the home side");
});

app.get("/", (req, res) => {
  res.send("Hello from the home side");
});

app.listen(PORT, () => {
  console.log(`Listening at the port ${PORT}`);
});
