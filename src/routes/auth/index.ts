import { doLogin, loginPage } from "controllers/auth";
import { Router } from "express";

const authRouter = Router();

authRouter.route("/login").get(loginPage).post(doLogin);

export default authRouter;
