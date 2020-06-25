import { Point, ApiResponse } from "../interfaces";

export async function getData(): Promise<ApiResponse[]> {
  const response = await fetch(`/api/`, {
    method: "GET",
    mode: "cors",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response.json();
}

export async function preparePointsToMap(
  data: Promise<ApiResponse[]>
): Promise<Point[]> {
  const points: Point[] = (await data).map((record) => ({
    vehicleType: record.VehicleType,
    lineNumber: record.LineNumber,
    latitude: record.Latitude,
    longitude: record.Longitude,
    vehicleNumber: record.VehicleNumber,
  }));

  return points;
}
