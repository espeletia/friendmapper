import { Link } from "react-router-dom";
import styles from "./Navbar.module.css";

export const Navbar = () => {
  return (
    <div className={styles.container}>
      <Link to="/">Home</Link>
      <Link to="/profile">Profile</Link>
      <Link to="/map">Map</Link>
    </div>
  );
};
