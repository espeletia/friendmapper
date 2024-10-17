import {
  useContext,
  createContext,
  useState,
  ReactNode,
  useMemo,
  useEffect,
  useRef,
} from "react";
import type { Marker, MarkerUpdateOptions } from "../../../types";

interface LeafletContextType {
  markers: Marker[];
  setMarkers: (markers: Marker[]) => void;
  updateMarkersDebounce: (
    options: MarkerUpdateOptions,
    debounce?: number,
  ) => void;
}

interface LeafletProviderProps {
  children: ReactNode;
}

const LeafletContext = createContext<LeafletContextType | undefined>(undefined);

export const useLeafletContext = () => {
  const context = useContext(LeafletContext);
  if (!context) {
    throw new Error("useLeafletContext must be used within a LeafletProvider");
  }
  return context;
};

export const LeafletProvider = (props: LeafletProviderProps) => {
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const [markers, setMarkers] = useState<Marker[]>([]);

  const updateMarkers = (options: MarkerUpdateOptions) => {
    console.log("in fetch options: ", options);
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((json) => {
        console.log("json", json);
        const newMarkers = json.map((item: any) => ({
          lat: Math.random() * 180 - 90,
          lng: Math.random() * 360 - 180,
          name: item.title,
        }));
        setMarkers(newMarkers);
      });
  };

  const updateMarkersDebounce = (
    options: MarkerUpdateOptions,
    debounce?: number,
  ) => {
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current);
    }
    timeoutRef.current = setTimeout(() => {
      updateMarkers(options);
    }, debounce ?? 1000);
  };

  useEffect(() => {
    console.log("in provider", markers);
  }, [markers]);

  const value = useMemo(
    () => ({
      markers,
      setMarkers,
      updateMarkersDebounce,
    }),
    [markers, setMarkers, updateMarkersDebounce],
  );

  return (
    <LeafletContext.Provider value={value}>
      {props.children}
    </LeafletContext.Provider>
  );
};
