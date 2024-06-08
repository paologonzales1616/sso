import path from 'path';

import express, { Express } from "express";

const app: Express = express();

app.set("views", path.join(__dirname + "/views"));
app.set("view engine", "ejs");

app.get("/healthz", (req, res) => res.status(200).send("OK"));

export default app;
