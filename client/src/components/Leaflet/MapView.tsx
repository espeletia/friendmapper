import React, { useEffect, useRef } from "react";
import L, { LatLngBoundsExpression } from "leaflet";
import { TileLayer, Marker, Popup, useMap, useMapEvents } from "react-leaflet";
import { Marker as MarkerType, MarkerUpdateOptions } from "../../types";

delete (L.Icon.Default.prototype as any)._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
});

interface Props {
  markers: MarkerType[] | null;
  updateMarkers: (options: MarkerUpdateOptions, debounce?: number) => void;
}

const MapView = (props: Props) => {
  const map = useMap();
  const isResizeMoveRef = useRef(false);

  useMapEvents({
    moveend: () => {
      if (isResizeMoveRef.current) {
        isResizeMoveRef.current = false;
        return;
      }

      const zoom = map.getZoom();
      const bounds = map.getBounds();

      props.updateMarkers({ bounds, zoom }, 2500);
      isResizeMoveRef.current = true;
    },
  });

  useEffect(() => {
    if (!props.markers) {
      return;
    }

    const bounds: LatLngBoundsExpression = props.markers.map((marker) => [
      marker.lat,
      marker.lng,
    ]);

    if (bounds?.length) {
      map.fitBounds(bounds);
    }
  }, [map]);

  return (
    <div>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {props?.markers?.map((marker, index) => (
        <Marker key={index} position={[marker.lat, marker.lng]}>
          <Popup>{marker.name}</Popup>
        </Marker>
      ))}
    </div>
  );
};

export default MapView;
