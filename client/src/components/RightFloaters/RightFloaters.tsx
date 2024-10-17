import { useRef, useState } from "react";
import { User } from "../../svg";
import Profile from "../Profile/Profile";
import styles from "./RightFloaters.module.css";

const RightFloaters = () => {
  const [userOpen, setUserOpen] = useState(false);
  const userButtonRef = useRef<HTMLButtonElement>(null);

  const handleUserOpen = () => {
    setUserOpen((prev) => !prev);
  };

  return (
    <div className={styles.container}>
      <button ref={userButtonRef} onClick={handleUserOpen}>
        <User />
      </button>
      {userOpen && <Profile closeCallback={handleUserOpen} />}
    </div>
  );
};

export default RightFloaters;
