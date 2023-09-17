import {
  Box,
  Center,
  Container,
  Heading,
  HStack,
  Image,
} from "@chakra-ui/react";

export const Navbar = () => {
  return (
    <Box w={"100%"} background={"white"} shadow={"md"} py={3}>
      <Center>
        <Container maxW={"container.xl"}>
          <HStack>
            <Image
              w={"80px"}
              src={
                "https://raw.githubusercontent.com/nimishbongale/DSC-RIT_Recruitment_Test/master/static/dscnew.png"
              }
            />
            <Heading color={"black"} size={"md"}>
              GDSC KMUTT
            </Heading>
          </HStack>
        </Container>
      </Center>
    </Box>
  );
};
