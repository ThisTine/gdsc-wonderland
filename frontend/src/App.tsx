import {
    Box,
    Button,
    Center,
    Container,
    Modal, ModalBody, ModalCloseButton,
    ModalContent, ModalFooter,
    ModalHeader,
    ModalOverlay, Text, useDisclosure,
} from "@chakra-ui/react";
import {Game} from "./Game.tsx";
import {NavBar} from "./Component/NavBar.tsx";

function App() {
    const {isOpen,onClose, onOpen} = useDisclosure()
    return (
        <>
        <Modal isOpen={isOpen} onClose={onClose} isCentered>
            <ModalOverlay />
            <ModalContent >
                <ModalHeader>GDSC KMUTT (ชมรมคนหาทำ) !!</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <Text>
                        GDSC (Google developer student clubs) เป็นโครงการที่มุ่งเน้นการสร้าง community ระหว่าง developer ภายในมหาลัยทั่วโลก 🌍 โดยมี "Google" บริษัทชั้นนำของโลกเป็นผู้สนับสนุน โดยจุดมุ่งหมายของเราคือการสร้างพื้นที่ เพื่อให้เหล่านักศึกษาภายในมหาลัยที่สนใจในด้านการพัฒนาเทคโนโลยี ได้เข้าพบปะ พูดคุย แลกเปลี่ยนความรู้ จนถึงการนำความรู้ไปประยุกต์ใช้ร่วมกันเพื่อก่อเกิดวัตกรรมใหม่ ๆ ให้กับชุมชน ✨
                    </Text>
                    <Text mt={5}>
                        โดยเกมนี้คาดหวังให้เราได้คุยกับเพื่อน ๆ ใหม่ ๆ และได้ลองทำกิจกรรมร่วมกัน !
                    </Text>
                </ModalBody>
                <ModalFooter/>
            </ModalContent>
        </Modal>
        <Box minW={"100vw"} minH={"100vh"}>
            <NavBar/>
            <Center>
                <Container maxW={"container.xl"}>
                    <Game/>
                </Container>
            </Center>

            <Button onClick={onOpen} pos={"fixed"} colorScheme={"blue"} rounded={"full"} bottom={5} right={3}>What is GDSC ?</Button>

        </Box>

        </>
    );
}

export default App
