import {Box, Center, Container, Heading, HStack, Image} from "@chakra-ui/react";

export const NavBar = () => {
	return (
		<Box w={"100%"} background={"white"} shadow={"md"} py={3}>
			<Center>
				<Container maxW={"container.xl"}>
					<HStack>
						<Image w={"80px"}
							   src={"https://raw.githubusercontent.com/nimishbongale/DSC-RIT_Recruitment_Test/master/static/dscnew.png?token=AKLHGWOKT4B6YUSLBUNLU526NEMIS"}/>
						<Heading color={"black"} size={"md"}>GDSC  KMUTT</Heading>
					</HStack>
				</Container>
			</Center>
		</Box>
	)
}