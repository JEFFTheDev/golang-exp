import React from "react";
import ReactDOM from "react-dom";
import { LatLng } from "leaflet";
import { MapContainer, TileLayer } from "react-leaflet";
import { Edge } from "../models/infra-models";

const defaultZoom = 13;
const defaultCenter = new LatLng(0, 0);

interface InfraMapProps {
  edges: Edge[];
}
export const InfraMap = ({ edges }: InfraMapProps) => {
  return (
    <MapContainer
      className="w-full h-full"
      center={defaultCenter}
      zoom={defaultZoom}
    >
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      {edges.map((edge) => {
        return <div>{edge}</div>;
      })}
    </MapContainer>
  );
};
