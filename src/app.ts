import express, { Express, Request, Response } from "express";

const app: Express = express();

app.get("/healthz", (req: Request, res: Response) => res.status(200));

export default app;
