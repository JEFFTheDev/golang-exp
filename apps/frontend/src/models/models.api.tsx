import { LatLng, LatLngExpression } from "leaflet";

export class TrainStation {
  constructor() {
    this.key = this.id;
  }

  id!: number;
  key!: number;
  name!: string;
  code!: string;
  lat!: number;
  lng!: number;
  connectedStations!: string[];

  getPosition(): LatLngExpression {
    return [this.lat, this.lng] as LatLngExpression;
  }
}
