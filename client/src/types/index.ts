import L from "leaflet";

export type MarkerType =
  | "KLUBY_FESTIVALY"
  | "KINA"
  | "PIVOVARY"
  | "DIVADLA_FILHARMONIE"
  | "MUZEA_GALERIE"
  | "PAMATKY"
  | "SPORT";

export interface Marker {
  Name: string;
  Description: string;
  Accessibility: boolean;
  AccessibilityNote: string;
  Capacity: number;
  CapacityNote: string;
  Phones: string;
  Email: string;
  Web: string;
  Okres: string;
  Obce: string;
  Address: string;
  lat: number;
  lng: number;
  Type: MarkerType;
}

export interface MarkerUpdateOptions {
  bounds: L.LatLngBounds;
  zoom: number;
}
