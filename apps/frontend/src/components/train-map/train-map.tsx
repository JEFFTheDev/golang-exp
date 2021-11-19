import { LatLng, LatLngExpression } from "leaflet";
import { useEffect, useState } from "react";
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  Polyline,
} from "react-leaflet";
import { TrainStation } from "../../models/models.api";
import { TrainApi } from "../../api/train.api";
import { find } from "lodash";

const position = [51.56215, 4.77258] as LatLngExpression;

export const TrainMap = () => {
  const [trainstations, setTrainStations] = useState<TrainStation[]>();
  const trainApi = new TrainApi();
  const drawnTracks: { [key: string]: string } = {};

  function trackIsDrawn(from: string, to: string): boolean {
    return (
      (drawnTracks[from] !== undefined ||
        Object.keys(drawnTracks).find((key) => drawnTracks[key] === from) !==
          undefined) &&
      (drawnTracks[to] !== undefined ||
        Object.keys(drawnTracks).find((key) => drawnTracks[key] === to) !==
          undefined)
    );
  }

  useEffect(() => {
    trainApi.getTrainStations().then((stations: TrainStation[]) => {
      setTrainStations(stations);
    });
  }, []);

  interface TrackAttributes {
    from: TrainStation;
    to: TrainStation;
  }
  const TrackToDraw = ({ from, to }: TrackAttributes) => {
    // TODO: do something about double tracks when reverse reference from destination station to source

    return (
      <Polyline
        key={from.id + to.id}
        positions={[from.getPosition(), to.getPosition()]}
      />
    );
  };

  return (
    <div style={{ width: "1000px", height: "1000px" }}>
      <MapContainer
        style={{ width: "1000px", height: "1000px" }}
        center={position}
        zoom={10}
        scrollWheelZoom={false}
      >
        <TileLayer
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        {trainstations?.map((i) => {
          return (
            <div>
              <Marker position={i.getPosition()}>
                <Popup>
                  {i.name} / {i.code}
                </Popup>
              </Marker>
              {i?.connectedStations?.map((t) => {
                return (
                  <TrackToDraw
                    from={i}
                    to={find(trainstations, { code: t }) as TrainStation}
                  />
                );
              })}
            </div>
          );
        })}
      </MapContainer>
    </div>
  );
};
