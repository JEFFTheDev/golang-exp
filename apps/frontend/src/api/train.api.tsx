import { render } from "@testing-library/react";
import axios, { AxiosInstance } from "axios";
import { plainToClass, Type } from "class-transformer";
import { TrainStation } from "../models/models.api";

const baseURL: string = "http://192.168.0.190:3001/api";

export class TrainApi {
  private readonly r: AxiosInstance;

  constructor() {
    this.r = axios.create({
      baseURL: baseURL,
    });
  }

  async getTrainStations(): Promise<TrainStation[]> {
    const res = await this.r.get<TrainStation[]>("/gettrainstations");
    return plainToClass(TrainStation, res.data);
  }
}
