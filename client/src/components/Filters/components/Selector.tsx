import styles from "./Selector.module.css";

interface Props {
  text: string;
}

const Selector = (props: Props) => {
  return (
    <div className={styles.container}>
      <p>{props.text}</p>
    </div>
  );
};

export default Selector;
