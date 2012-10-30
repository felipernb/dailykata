package com.feliperibeiro.bowling;

public class BowlingGame {
	int[] rolls;
	int currentRoll;

	public BowlingGame() {
		this.rolls = new int[21];
	}

	public void roll(int p) {
		rolls[currentRoll++] = p;
	}

	public int score() {
		int score = 0;
		int frame = 0;

		for (int i = 0; i < 10; i++) {
			if (isStrike(frame)) {
				score += 10 + strikeBonus(frame);
				frame++;
			} else if (isSpare(frame)) {
				score += 10 + spareBonus(frame);
				frame += 2;
			} else {
				score += sumOfRolls(frame);
				frame += 2;
			}
		}

		return score;
	}

	private boolean isStrike(int frame) {
		return rolls[frame] == 10;
	}

	private boolean isSpare(int frame) {
		return sumOfRolls(frame) == 10;
	}

	private int strikeBonus(int frame) {
		return sumOfRolls(frame+1);
	}

	private int spareBonus(int frame) {
		return rolls[frame+2];
	}

	private int sumOfRolls(int frame) {
		return rolls[frame] + rolls[frame+1];
	}
}
