import {
  useContext,
  createContext,
  useState,
  ReactNode,
  useMemo,
  useEffect,
} from "react";
import type { Marker } from "../../../types";

interface LeafletContextType {
  markers: Marker[];
  setMarkers: (markers: Marker[]) => void;
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
  const [markers, setMarkers] = useState<Marker[]>([]);

  useEffect(() => {
    console.log("in provider", markers);
  }, [markers]);
  const value = useMemo(() => ({ markers, setMarkers }), [markers, setMarkers]);

  return (
    <LeafletContext.Provider value={value}>
      {props.children}
    </LeafletContext.Provider>
  );
};
