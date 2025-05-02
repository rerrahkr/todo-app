/**
 * Generator of sequential numbers starting from 0.
 *
 * @param end If provided, the generator will yield numbers from 0 to end - 1.
 *            If not provided, the generator will yield numbers indefinitely.
 * @returns A generator that yields sequential numbers.
 */
export function* seq(end?: number): Generator<number> {
  let i = 0;
  while (end === undefined || i < end) {
    yield i++;
  }
}
