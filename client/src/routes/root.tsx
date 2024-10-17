import { Outlet } from "react-router-dom";
import Navbar from "../components/ui/Navbar";

import styles from "./root.module.css";

const Root = () => {
  return (
    <div className={styles.container}>
      <Navbar />
      <p>layout</p>
      <Outlet />
    </div>
  );
};

export default Root;
