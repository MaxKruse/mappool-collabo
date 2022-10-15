export interface Slot {
    name: string;
    index: number;
}

export interface DifficultyAttributes {
    hp: number;
    od: number;
    ar: number;
    cs: number;
    stars: number;
    modStrings: string[];
    modInts?: number;
}

export interface Map {
    id: number;
    name: string;
    slot: Slot;
    description: string;
    difficultyAttributes: DifficultyAttributes;
}

export interface Vote {
    id: number;
    value: number;
    comment: string;
    author: Author;
}

export interface Suggestion {
    id: number;
    comment: string;
    map: Map;
    voteScore: number;
    votes: Vote[];
}

export interface Round {
    id: number;
    name: string;
    suggestions: Suggestion[];
}

export interface Tournament {
    id: number;
    name: string;
    description: string;
    owner: Owner;
    poolers: Pooler[];
    testplayers: Testplayer[];
    rounds: Round[];
}