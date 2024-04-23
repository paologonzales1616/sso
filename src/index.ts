import dotenv from "dotenv";
import server from "./app";

dotenv.config();

const PORT = process.env["PORT"] || 3000;

server.listen(PORT, () => console.log(`Server running at localhost:${PORT}`));
