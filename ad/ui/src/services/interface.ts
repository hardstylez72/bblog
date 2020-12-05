abstract class Service<T> {
  abstract GetList(): Promise<T[]>
  abstract Delete(id: number): Promise<void>
  abstract Create(t: T): Promise<T>
  abstract Update(t: T): Promise<T>
}
