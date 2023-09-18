export type InitialResponse = {
	sessionId: string,
	email: string,
	pictures: InitialPicture[],
}

export type InitialPicture = {
	id: string,
	src: string,
	title: string,
	description: string
}

export type CommitResponse = {
	matched: boolean,
	forwardLink: string,
	pairedWith: string,
	pairedDiff: number,
};
