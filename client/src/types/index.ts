import L from "leaflet";

export interface Marker {
  name: string;
  lat: number;
  lng: number;
}

export interface MarkerUpdateOptions {
  bounds: L.LatLngBounds;
  zoom: number;
}
