import {
    Box,
    Button,
    Center, CircularProgress,
    Container, Flex, Heading,
    HStack,
    Link,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    Text,
    useDisclosure,
} from "@chakra-ui/react";
import {Game} from "./Game.tsx";
import {Navbar} from "./components/Navbar.tsx";
import bg from "./assets/bg.jpg";
import {useContext} from "react";
import {GlobalContext} from "./contexts/GlobalState.tsx";

function App() {
    const {state, error} = useContext(GlobalContext);
    const {isOpen, onClose, onOpen} = useDisclosure();

    return (
        <>
            <Modal isOpen={isOpen} onClose={onClose} isCentered>
                <ModalOverlay/>
                <ModalContent>
                    <ModalHeader>GDSC KMUTT (ชมรมคนหาทำ) !!</ModalHeader>
                    <ModalCloseButton/>
                    <ModalBody>
                        <Text>
                            GDSC (Google developer student clubs)
                            เป็นโครงการที่มุ่งเน้นการสร้าง community ระหว่าง developer
                            ภายในมหาลัยทั่วโลก 🌍 โดยมี "Google"
                            บริษัทชั้นนำของโลกเป็นผู้สนับสนุน
                            โดยจุดมุ่งหมายของเราคือการสร้างพื้นที่
                            เพื่อให้เหล่านักศึกษาภายในมหาลัยที่สนใจในด้านการพัฒนาเทคโนโลยี
                            ได้เข้าพบปะ พูดคุย แลกเปลี่ยนความรู้
                            จนถึงการนำความรู้ไปประยุกต์ใช้ร่วมกันเพื่อก่อเกิดนวัตกรรมใหม่ ๆ
                            ให้กับชุมชน ✨
                        </Text>
                        <Text mt={5}>
                            โดยเกมนี้คาดหวังให้เราได้คุยกับเพื่อน ๆ ใหม่ ๆ
                            และได้ลองทำกิจกรรมร่วมกัน !
                        </Text>
                        <HStack>
                            <Button colorScheme="facebook" as={"a"}>
                                <Link
                                    href={"https://www.facebook.com/gdsc.kmutt"}
                                    target={"_blank"}
                                >
                                    Facebook
                                </Link>
                            </Button>
                            <Button colorScheme="pink" as={"a"}>
                                <Link
                                    href={"https://www.instagram.com/gdsc.kmutt"}
                                    target={"_blank"}
                                >
                                    Instagram
                                </Link>
                            </Button>
                        </HStack>
                    </ModalBody>
                    <ModalFooter/>
                </ModalContent>
            </Modal>
            <Box
                color={"white"}
                minW={"100vw"}
                backdropFilter={"blur(10px)"}
                bgImage={bg}
                backgroundSize={"cover"}
                minH={"100vh"}
            >
                <Navbar/>
                <Center>
                    <Container maxW={"container.xl"} height="100vh">
                        {state ? (
                            <Game/>
                        ) : error !== null ? (
                            <Flex alignItems="center" justifyItems="center"
                                  style={{width: '100%', height: '100%'}}>
                                <Heading textAlign={"center"}>Loading</Heading>
                                <CircularProgress isIndeterminate/>
                            </Flex>
                        ) : (
                            <Flex alignItems="center" justifyItems="center"
                                  style={{width: '100%', height: '100%'}}>
                                <Heading textAlign={"center"}>
                                    Token not found, please navigate from Google Form confirmation email!
                                </Heading>
                            </Flex>
                        )}
                    </Container>
                </Center>

                <Button
                    onClick={onOpen}
                    pos={"fixed"}
                    colorScheme={"blue"}
                    rounded={"full"}
                    bottom={5}
                    right={3}
                >
                    What is GDSC ?
                </Button>
            </Box>
        </>
    );
}

export default App;
