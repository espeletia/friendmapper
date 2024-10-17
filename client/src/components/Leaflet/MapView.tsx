import React, { useEffect } from "react";
import L, { LatLngBoundsExpression } from "leaflet";
import { TileLayer, Marker, Popup, useMap } from "react-leaflet";
import { Marker as MarkerType } from "../../types";

delete (L.Icon.Default.prototype as any)._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
});

interface Props {
  markers: MarkerType[];
}

const MapView = (props: Props) => {
  const map = useMap();

  useEffect(() => {
    if (props.markers.length) {
      const bounds: LatLngBoundsExpression = props.markers.map((marker) => [
        marker.lat,
        marker.lng,
      ]);
      map.fitBounds(bounds);
    }
  }, [props.markers, map]);

  return (
    <div>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {props.markers.map((marker, index) => (
        <Marker key={index} position={[marker.lat, marker.lng]}>
          <Popup>{marker.name}</Popup>
        </Marker>
      ))}
    </div>
  );
};

export default MapView;
