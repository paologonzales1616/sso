import express, { Express, Request, Response } from "express";

const app: Express = express();

app.get("/healthz", (req, res) => res.status(200).send('OK'));

export default app;
