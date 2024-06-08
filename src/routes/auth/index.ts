import { Router } from "express";

const authRouter = Router();

authRouter
  .get("/login", (req, res) => res.render('pages/login'))
  .post("/login", (req, res) => res.status(200).end());

export default authRouter;
