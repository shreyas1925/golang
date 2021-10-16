const express = require("express");
const app = express();
const PORT = process.env.PORT || 5000;

app.get("/", (req, res) => {
  res.send("Hello from the home side");
});

app.listen(PORT, () => {
  console.log(`Listening at the port ${PORT}`);
});
