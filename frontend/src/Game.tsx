import {
  Box,
  Center,
  Heading,
  Image,
  VStack,
  useToast,
} from "@chakra-ui/react";
import axios, { AxiosError } from "axios";
import { useEffect, useLayoutEffect, useState } from "react";
import hacker from "./assets/tine_warning.jpg";
import seedRandom from "seedrandom";
import { ErrorResponse, SuccessResponse } from "./types/payload.ts";

const commit = async (pic: string, session: string) => {
  const { data } = await axios.post("/wonderland/api/pair/commit", {
    sessionNo: session,
    itemNo: pic,
  });
  return data;
};

const GameButton = ({ pic, onClick }: { pic: string; onClick: () => void }) => {
  return (
    <Box
      display={"flex"}
      justifyContent={"center"}
      alignItems={"center"}
      w={"350px"}
      height={"350px"}
      overflow={"hidden"}
      onClick={onClick}
      mt={5}
      cursor="pointer"
      _hover={{ transform: "scale(1.1)" }}
      transition={"all 0.1s"}
      p={3}
      bg={"white"}
      shadow={"md"}
      rounded={"xl"}
    >
      <Image maxW={"300px"} src={pic} />
    </Box>
  );
};

const Loading = ({ pic, hack }: { pic: string; hack: boolean }) => {
  return (
    <Box
      display={"flex"}
      justifyContent={"center"}
      alignItems={"center"}
      w={"350px"}
      height={"350px"}
      overflow={"hidden"}
      mt={5}
      cursor="pointer"
      _hover={{ transform: "scale(1.1)" }}
      transition={"all 0.1s"}
      p={3}
      bg={"green"}
      shadow={"md"}
      rounded={"xl"}
    >
      <Image maxW={"300px"} src={hack ? hacker : pic} />
    </Box>
  );
};

export const Game = ({ token, unhash }: { token: string; unhash: string }) => {
  const [pic, setPic] = useState<string>("");
  const [timeOutLimit, setTimeOutLimit] = useState<number>(0);
  const [Loadtime, setLoadtime] = useState<string>("");
  const [YouareHacker, setYouareHacker] = useState<number>(0);
  const toast = useToast();
  const rnd = seedRandom(token, { entropy: true });
  useLayoutEffect(() => {
    const timer = setInterval(() => {
      setPic(pictures[Math.floor(rnd() * pictures.length)]);
    }, (Math.random() * 0.5 + 0.5) * 1000);
    return () => clearInterval(timer);
  }, [rnd]);
  useEffect(() => {
    if (timeOutLimit) {
      setTimeout(() => {
        setTimeOutLimit((v) => v - 1);
      }, 1000);
    }
  }, [timeOutLimit]);

  const onClick = () => {
    setLoadtime(pic!);
    commit(pictures.indexOf(pic).toString(), unhash)
      .then((res: SuccessResponse) => {
        if (res.data.matched) {
          window.location.replace(res.data.forwardLink);
        } else {
          setTimeOutLimit(5);
          setLoadtime("");
        }
      })
      .catch((res: AxiosError<ErrorResponse>) => {
        setLoadtime("");
        toast({
          status: "error",
          title: res.response?.data.message || res.message,
        });
        setYouareHacker(5);
      });
  };
  return (
    <Box pt={5}>
      <Center>
        <VStack>
          <Heading>Find some friends, </Heading>
          <Heading size={"lg"} textAlign={"center"}>
            then tap the picture together !
          </Heading>
          {timeOutLimit ? (
            <Box mt={10}>
              <Heading textAlign={"center"}>Tap at the same TIME!!</Heading>
              <Heading textAlign={"center"}>
                Let's try again in {timeOutLimit}{" "}
              </Heading>
            </Box>
          ) : Loadtime ? (
            <Loading pic={Loadtime} hack={false} />
          ) : YouareHacker ? (
            <Loading pic={Loadtime} hack={true} />
          ) : (
            <GameButton onClick={onClick} pic={pic} />
          )}
        </VStack>
      </Center>
    </Box>
  );
};
