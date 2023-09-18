import {
	Box,
	Center,
	Heading,
	Image,
	VStack,
	useToast, Alert, Flex, Text, Button,
} from '@chakra-ui/react'
import axios, { AxiosError } from 'axios'
import { useContext, useEffect, useLayoutEffect, useState } from 'react'
import hacker from './assets/tine_warning.jpg'
import seedRandom from 'seedrandom'
import { GlobalContext } from './contexts/GlobalState.tsx'
import { ErrorResponse, SuccessResponse } from './types/response.ts'
import { CommitResponse, InitialPicture } from './types/payload.ts'

const commit = async (picId: string, session: string) => {
	const { data } = await axios.post('/wonderland/api/pair/commit', {
		sessionId: session,
		itemId: picId,
	})
	return data
}

const GameButton = ({ pic, onClick }: { pic: string; onClick: () => void }) => {
	return (
		<Box
			display={'flex'}
			justifyContent={'center'}
			alignItems={'center'}
			w={'350px'}
			height={'350px'}
			overflow={'hidden'}
			onClick={onClick}
			mt={5}
			cursor='pointer'
			_hover={{ transform: 'scale(1.1)' }}
			transition={'all 0.1s'}
			p={3}
			bg={'white'}
			shadow={'md'}
			rounded={'xl'}
		>
			<Image maxW={'300px'} src={pic} />
		</Box>
	)
}

const Loading = ({ pic, hack }: { pic: string; hack: boolean }) => {
	return (
		<Box
			display={'flex'}
			justifyContent={'center'}
			alignItems={'center'}
			w={'350px'}
			height={'350px'}
			overflow={'hidden'}
			mt={5}
			cursor='pointer'
			_hover={{ transform: 'scale(1.1)' }}
			transition={'all 0.1s'}
			p={3}
			bg={'green'}
			shadow={'md'}
			rounded={'xl'}
		>
			<Image maxW={'300px'} src={hack ? hacker : pic} />
		</Box>
	)
}

const calculateReward = (diff: number) => {
	if (diff < 1000) {
		return 'Oreo'
	}
	if (diff < 2000) {
		return 'Pocky'
	}
	return 'Pi-Po Yelly'
}

export const Game = () => {
	const toast = useToast()
	const { state } = useContext(GlobalContext)
	const [picture, setPicture] = useState<InitialPicture>({
		id: '',
		src: '',
		title: '',
		description: '',
	})
	const [lockedPicture, setLockedPicture] = useState<string | null>(null)
	const [hackerPicture, setHackerPicture] = useState<number>(0)
	const [timeoutCountdown, setTimeoutCountdown] = useState<number>(0)
	const [paired, setPaired] = useState<CommitResponse | null>(null,
	)
	const rnd = seedRandom(state!.sessionId, { entropy: true })

	useLayoutEffect(() => {
		setPicture(state!.pictures[Math.floor(rnd() * state!.pictures.length)])
	}, [])

	useLayoutEffect(() => {
		const timer = setInterval(() => {
			if (lockedPicture || timeoutCountdown > 0) return
			setPicture(state!.pictures[Math.floor(rnd() * state!.pictures.length)])
		}, (Math.random() * 0.5 + 0.5) * 1000)
		return () => clearInterval(timer)
	}, [rnd])

	useEffect(() => {
			if (timeoutCountdown) {
				setTimeout(() => {
						if (timeoutCountdown === 1) {
							setPicture(state!.pictures[Math.floor(rnd() * state!.pictures.length)])
						}
						setTimeoutCountdown((v) => v - 1)
					},
					1000,
				)
			}
		}, [timeoutCountdown],
	)

	const onClick = () => {
		setLockedPicture(picture.src)
		commit(picture.id, state!.sessionId)
			.then((res: SuccessResponse<CommitResponse>) => {
				if (res.data.matched) {
					setPaired(res.data)
				} else {
					setTimeoutCountdown(5)
					setLockedPicture(null)
				}
			})
			.catch((res: AxiosError<ErrorResponse>) => {
				setLockedPicture(null)
				toast({
					status: 'error',
					title: res.response?.data.message || res.message,
				})
				setHackerPicture(5)
			})
	}
	return (
		<Box pt={5}>
			<Center>
				<VStack>
					{
						!paired && (
							<>
								<Heading>Find some friends,</Heading>
								<Heading size={'lg'} textAlign={'center'}>
									then tap the same picture together at the same time!
								</Heading></>
						)
					}

					{timeoutCountdown ? (
						<Box mt={10}>
							<Heading textAlign={'center'}>No matching found,<br />tap at the same TIME!!</Heading>
							<Heading textAlign={'center'}>
								Let's try again in {timeoutCountdown}
							</Heading>
							<Alert status='info' style={{ borderRadius: 16 }} mt={16}>
								<Flex direction='column'>
									<img src={picture.src} width='120px' />
									<Text color='#3a3a3a' mt={6}>
										<b>{picture.title}</b>
										<p>{picture.description}</p>
									</Text>
								</Flex>
							</Alert>
						</Box>
					) : paired ? (
						<Box mt={10} position='relative'>
							<Heading textAlign={'center'}>Congratulation 🎉</Heading>
							<Box position='absolute'
								 top={24}
								 left={-12}
								 width={250}
								 height={250}
								 borderRadius={125}
								 backgroundColor='#435287'
								 display='flex'
								 flexDirection='column'
								 justifyContent='center'
								 alignItems='center'>
								<Text>Paired with </Text>
								<Heading textAlign={'center'} size='md' mb={4}>{paired.pairedWith}</Heading>
							</Box>
							<Box position='absolute'
								 top={271}
								 left={32}
								 width={180}
								 height={180}
								 borderRadius={90}
								 backgroundColor='#89AC76'
								 display='flex'
								 flexDirection='column'
								 justifyContent='center'
								 alignItems='center'>
								<Text>Delay </Text>
								<Heading textAlign={'center'} size='md'>
									{Math.round(paired.pairedDiff / 1000)} secs<br />
									{paired.pairedDiff % 1000} ms</Heading>
							</Box>
							<Box position='absolute'
								 top={295}
								 left={0}
								 width={150}
								 height={150}
								 borderRadius={75}
								 backgroundColor='#CB2821'
								 display='flex'
								 flexDirection='column'
								 justifyContent='center'
								 alignItems='center'>
								<Text>Reward </Text>
								<Heading textAlign={'center'} size='md'>
									{calculateReward(paired.pairedDiff)}</Heading>
							</Box>
							<Button
								onClick={() => {
									window.location.href = paired.forwardLink
								}}
								pos={'absolute'}
								colorScheme={'green'}
								rounded={'full'}
								transform={'translate(-50%, -50%)'}
								top={490}
								right={0}
							>
								Go to activity forms
							</Button>
						</Box>
					) : lockedPicture ? (
						<Loading pic={lockedPicture} hack={false} />
					) : hackerPicture ? (
						<Loading pic={lockedPicture!} hack={true} />
					) : (
						<GameButton onClick={onClick} pic={picture.src} />
					)}
				</VStack>
			</Center>
		</Box>
	)
}
