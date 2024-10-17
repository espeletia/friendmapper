import Filters from "../components/Filters/Filters";
import Leaflet from "../components/Leaflet";

import styles from "./map.module.css";

const MapRoute = () => {
  return (
    <div className={styles.container}>
      <Filters />
      <Leaflet />
    </div>
  );
};

export default MapRoute;
