import path from 'path';

import express, { Express } from "express";
import authRouter from "routes/auth";

const app: Express = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.set("views", path.join(__dirname + "/views"));
app.set("view engine", "ejs");

app.get("/healthz", (req, res) => res.status(200).send("OK"));

app.use("/auth", authRouter);

export default app;
