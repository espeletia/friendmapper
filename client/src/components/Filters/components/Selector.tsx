import cx from "classnames";
import { ReactNode } from "react";

import styles from "./Selector.module.css";

interface Props {
  children: ReactNode;
  index: number;
  handleChange: (index: number) => void;
  selected?: boolean;
}

const Selector = (props: Props) => {
  return (
    <button
      onClick={() => props.handleChange(props.index)}
      className={cx(styles.container, { [styles.selected]: props.selected })}
    >
      {props.children}
    </button>
  );
};

export default Selector;
