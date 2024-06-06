import dotenv from "dotenv";
import server from "./app";

dotenv.config();

const PORT = process.env["PORT"] || 3000;

const start = () => {
  try {
    server.listen(PORT, () =>
      console.log(`Server running at localhost:${PORT}`)
    );
  } catch (error) {
    console.error(error);
    process.exit(1);
  }
};

void start();
