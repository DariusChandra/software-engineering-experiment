import { add, subtract } from '../src/math';

describe('Math functions', () => {
    test('Addition', () => {
        expect(add(1, 2)).toBe(3);
    });

    test('Subtraction', () => {
        expect(subtract(5, 3)).toBe(2);
    });
});
