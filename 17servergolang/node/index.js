const express = require("express");
const app = express();
const PORT = process.env.PORT || 5000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

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

app.get("/get", (req, res) => {
  res.json({ message: "Hello going to build mini app in Golang " });
});

app.post("/post", (req, res) => {
  let mydata = req.body;
  res.status(200).send(mydata);
});

// using post by passing data through form
app.post("/postformdata", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
  console.log(`Listening at the port ${PORT}`);
});
